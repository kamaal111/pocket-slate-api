import sys

if __name__ == "__main__":
    arguments = sys.argv[1:]
    assert len(arguments) == 2
    assert arguments[0] == arguments[1]
