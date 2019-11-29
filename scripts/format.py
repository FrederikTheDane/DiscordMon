import subprocess


def format_code():
    subprocess.run(["go", "fmt", "github.com/frederikthedane/DiscordMon/..."])


if __name__ == '__main__':
    format_code()
