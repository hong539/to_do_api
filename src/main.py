from flask import Flask, request, send_file
import qrcode
import io

app = Flask(__name__)

@app.route('/qrcode', methods=['POST'])
def generate_qrcode():
    data = request.get_json()
    qr = qrcode.QRCode(version=1, box_size=10, border=4)
    qr.add_data(data['text'])
    qr.make(fit=True)
    img = qr.make_image(fill_color=data['color'], back_color=data['bgcolor'])
    img_io = io.BytesIO()
    img.save(img_io, 'JPEG', quality=70)
    img_io.seek(0)
    return send_file(img_io, mimetype='image/jpeg')

if __name__ == '__main__':
    app.run(debug=True)