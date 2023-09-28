while True:
    try:
        n = int(input())
    except EOFError:
        break

    out = "vai ter copa!"
    if n > 0:
        out = "vai ter duas!"
    print(out)

