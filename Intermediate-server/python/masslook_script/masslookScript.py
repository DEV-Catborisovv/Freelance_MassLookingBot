from telethon import TelegramClient, types, functions
from telethon.errors import SessionPasswordNeededError, FloodWaitError
from telethon.sessions import StringSession

import sys
import asyncio
import time

session_string = ""

async def main(api_id, api_hash, phone_number, listofchats):
    client = TelegramClient(StringSession(session_string), api_id, api_hash)
    
    await client.connect()

    if not await client.is_user_authorized():
        try:
            await client.send_code_request(phone_number)
            print('Enter the code you received: ', flush=True)
            code = sys.stdin.readline().strip()
            try:
                await client.sign_in(phone_number, code)
                
                print(code)
            except SessionPasswordNeededError:
                print('Your 2FA password: ', flush=True)
                password = sys.stdin.readline().strip()
                await client.sign_in(password=password)
        except Exception as e:
            print(f"Failed to authenticate: {e}")
            await client.disconnect()            
            return

    # Convert list of chats to a set for fast lookup
    chat_set = set(listofchats)

    try:
        async for dialog in client.iter_dialogs(limit=500):
            if dialog.is_group and dialog.name in chat_set:
                print(f"Processing chat: {dialog.name}")
                async for user in client.iter_participants(entity=dialog.entity, limit=1000):
                    if (
                        not user.stories_unavailable and 
                        not user.stories_hidden and 
                        user.stories_max_id
                    ):
                        try:
                            res = await client(
                                functions.stories.ReadStoriesRequest(
                                    peer=user.id, 
                                    max_id=user.stories_max_id
                                )
                            )
                            print(f"Viewed stories of user: {user.id} in chat: {dialog.name}")
                            time.sleep(5)  # Sleep to avoid rate limiting
                        except FloodWaitError as e:
                            print(f"Rate limited. Sleeping for {e.seconds} seconds.")
                            time.sleep(e.seconds)
                        except Exception as e:
                            print(f"Failed to view stories for user {user.id}: {e}")
    except Exception as e:
        print(f"An error occurred: {e}")

    # Send a message to "Saved Messages" upon completion
    try:
        await client.send_message('me', 'Script has completed successfully.')
        print("Sent completion message to Saved Messages.")
    except Exception as e:
        print(f"Failed to send completion message: {e}")

    await client.disconnect()

if __name__ == "__main__":
    if len(sys.argv) < 5:
        print("Usage: python script.py <api_id> <api_hash> <phone_number> <chat1> <chat2> ...")
        sys.exit(1)
    
    api_id = int(sys.argv[1])
    api_hash = str(sys.argv[2])
    phone_number = str(sys.argv[3])
    listofchats = sys.argv[4:]

    asyncio.run(main(api_id, api_hash, phone_number, listofchats))
