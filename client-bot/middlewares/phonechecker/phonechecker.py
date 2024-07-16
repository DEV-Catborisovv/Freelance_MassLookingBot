import re

def is_valid_phone_number(phone_number):
    # Регулярное выражение для проверки формата телефонного номера
    pattern = re.compile(r'^\+\d{1,3}\d+$')
    
    # Проверка соответствия строки регулярному выражению
    if pattern.match(phone_number):
        return True
    else:
        return False