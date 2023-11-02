import os
import requests
from dotenv import load_dotenv

load_dotenv()  # Load environment variables from .env file

PROGRAMMERBAR_API_TOKEN = os.getenv('PROGRAMMERBAR_API_TOKEN')
if not PROGRAMMERBAR_API_TOKEN:
    print("PROGRAMMERBAR_API_TOKEN is not set")
    exit(1)

if len(os.sys.argv) < 2:
    print(f"Usage: {os.sys.argv[0]} <new status>")
    exit(1)

NEW_STATUS = os.sys.argv[1]

try :
    NEW_STATUS = int(NEW_STATUS)
except ValueError:
    print(f"Status must be an integer. Got {NEW_STATUS}")
    exit(1)


url = 'https://programmerbar.fly.dev/status'
headers = {
    'Authorization': f'Bearer {PROGRAMMERBAR_API_TOKEN}',
    'Content-Type': 'application/json',
}
data = {
    'status': NEW_STATUS
}

response = requests.post(url, headers=headers, json=data)

if response.status_code == 200:
    print("Status successfully updated.")
else:
    print(f"Failed to update status. Status code: {response.status_code}")
    print(response.text)
