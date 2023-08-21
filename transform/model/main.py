from http.server import ThreadingHTTPServer, BaseHTTPRequestHandler
import json
from resnet50 import Resnet50
 

host = ('0.0.0.0', 8888)
resnet50=Resnet50()

class Resquest(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(400,"should use post method")
        self.send_header('Content-type', 'application/json')
        self.end_headers()

    def do_POST(self):
        self.send_response(200)
        self.send_header('Content-type', 'application/json')
        self.end_headers()
        resq_data={'Vector':{},'Msg':""}
        try:
            req_data = self.rfile.read(int(self.headers["Content-Length"])).decode()
            req_body=json.loads(req_data)

            vector = resnet50.toVector(req_body["ImgPath"])
            resq_data["Vector"]=vector
        except Exception as e:
            resq_data["Msg"]=repr(e)
        finally:
            self.wfile.write(json.dumps(resq_data).encode())


if __name__ == '__main__':
    server = ThreadingHTTPServer(host, Resquest)
    print("Starting server, listen at: %s:%s" % host)
    server.serve_forever()