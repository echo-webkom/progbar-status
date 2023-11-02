import requests

url = 'https://programmerbar.fly.dev/status'

response = requests.get(url)

if response.status_code == 200:
    print(f"Status: {response.json()}")
else:
    print(f"Failed to fetch status. Status code: {response.status_code}")
    print(response.text)
