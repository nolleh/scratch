import subprocess
import re

def main():
    battery = subprocess.check_output(['bash', '-c', 'pmset -g batt'])
    battery = battery.decode('utf-8')

    pattern = re.compile(r"\d+%")
    matches = pattern.findall(battery)
    if len(matches) == 0:
        return []
    percent = matches[0].replace('%', '')
    print(f'percent: {percent}')

    pattern = re.compile(r"; .*;")
    matches = pattern.findall(battery)
    status = matches[0].replace(';', '').strip() if len(matches) > 0 else ''
    print(f'status: {status}')
main()
