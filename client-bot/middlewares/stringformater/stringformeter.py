def format_string(input_string:str) -> tuple[list[str], bool]:
    # Разделяем строку по запятым
    elements = input_string.split(',')
    
    elements = [element.strip() for element in elements]
    
    if all(elements):
        return elements, True
    else:
        return [], False

