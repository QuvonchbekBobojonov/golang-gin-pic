import http.client

conn = http.client.HTTPSConnection("127.0.0.1", 8080)
payload = ''
headers = {
  'Authorization': 'Basic bW9vcmZvIDEyMzQ='
}
conn.request("GET", "/videos", payload, headers)
res = conn.getresponse()
data = res.read()
print(data.decode("utf-8"))