from telethon import TelegramClient, types, functions
from telethon.errors import SessionPasswordNeededError
import sys
import asyncio

async def main(api_id, api_hash, phone_number, listofchats):
    client = TelegramClient(f"./python/masslook_script/sessions/{phone_number}.session", api_id, api_hash)
    
    await client.connect()

    if not await client.is_user_authorized():
        try:
            await client.send_code_request(phone_number)
            code = input('Enter the code you received: ')
            try:
                await client.sign_in(phone_number, code)
            except SessionPasswordNeededError:
                password = input('Your 2FA password: ')
                await client.sign_in(password=password)
        except Exception as e:
            print(f"Failed to authenticate: {e}")
            await client.disconnect()
            return

    # Преобразуем список чатов в множество для быстрого поиска
    chat_set = set(listofchats)

    async for dialog in client.iter_dialogs(limit=500):
        if dialog.is_group and dialog.name in chat_set:
            print(f"Processing chat: {dialog.name}")
            async for user in client.iter_participants(entity=dialog.entity, limit=1000):
                if (
                    not user.stories_unavailable and 
                    not user.stories_hidden and 
                    user.stories_max_id
                ):
                    res = await client(
                        functions.stories.ReadStoriesRequest(
                            peer=user.id, 
                            max_id=user.stories_max_id
                        )
                    )
                    print(f"Viewed stories of user: {user.id} in chat: {dialog.name}")

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