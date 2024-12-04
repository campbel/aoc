def load_input(file_path: str) -> str:
    with open(file_path, "r") as f:
        return f.read()


def scan_int(cursor: int, input_data: str) -> tuple[int, int]:
    digits = ""
    while cursor < len(input_data) and input_data[cursor].isdigit():
        digits += input_data[cursor]
        cursor += 1
    return (int(digits) if digits else 0, cursor)


def scan_params(cursor: int, input_data: str) -> tuple[tuple[int, int], int, bool]:
    if input_data[cursor] != "(":
        return ((0, 0), cursor, False)
    a, cursor = scan_int(cursor + 1, input_data)
    if input_data[cursor] != ",":
        return ((a, 0), cursor, False)
    b, cursor = scan_int(cursor + 1, input_data)
    if input_data[cursor] != ")":
        return ((a, b), cursor, False)
    return ((a, b), cursor, True)


def scan_for_do(cursor: int, input_data: str) -> int:
    while cursor < len(input_data) and input_data[cursor : cursor + 4] != "do()":
        cursor += 1
    return cursor


def main():
    input_data = load_input("input.txt")

    # List of operands to multiply and add
    operands: list[tuple[int, int]] = []

    cursor = 0
    while cursor < len(input_data):
        if input_data[cursor : cursor + 5] == "don't":
            cursor = scan_for_do(cursor + 5, input_data)
            continue
        if input_data[cursor : cursor + 3] == "mul":
            params, cursor, valid = scan_params(cursor + 3, input_data)
            if not valid:
                continue
            operands.append(params)
        cursor += 1

    sum_of_products = 0
    for a, b in operands:
        sum_of_products += a * b

    print(sum_of_products)


if __name__ == "__main__":
    main()
