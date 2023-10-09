def duplicate_count(text: str) -> int:
    
    occurances: dict[str|int, int] = {}
    num_duplicates: int = 0
    
    for char in text:
        char = char.lower()
        if char not in occurances.keys():
            occurances[char] = 1
        else:
            occurances[char] += 1
    
    for value, count in occurances.items():
        if count >= 2:
            num_duplicates += 1
    
    return num_duplicates
     