def alphabet_position(text: str) -> str:
    
    alphabet: str = "abcdefghijklmnopqrstuvwxyz"
    new_string: list[str] = []
    
    for char in text:
        char: str = char.lower()
        if char not in alphabet:
            continue
        position: str = str(alphabet.index(char) + 1)
        new_string.append(position)
    
    return " ".join(new_string)