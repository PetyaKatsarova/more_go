learning from https://github.com/icoder-new/goBomber

goBomber is a simple SMS bombing tool written in Go. It allows you to send multiple SMS messages to a specific phone number using various service providers. (THIS IS BETA VERSION...)

Installation
Make sure you have Go installed on your system. You can download it from the official website: https://golang.org/
Clone the goBomber repository:
git clone https://github.com/icoder-new/goBomber.git
Navigate to the project directory:
cd goBomber
Build the project:
go build
Usage
To use goBomber, run the executable and provide the phone number you want to target as a command-line argument:

./goBomber -number=<phone-number-without-992>
Replace <phone-number> with the actual phone number you want to bombard with SMS messages.

Supported Service Providers
goBomber currently supports the following service providers:

Somon
Avrang
Dastras
When you run goBomber, it will simultaneously send SMS messages to the target phone number using these service providers. The tool will display the status code received from each service provider.

Disclaimer
Please use goBomber responsibly and only on phone numbers that you have permission to target. Sending unsolicited SMS messages or using goBomber for malicious purposes is illegal and unethical.

The author of goBomber is not responsible for any misuse or damage caused by this tool. Use it at your own risk.

License
This project is licensed under the MIT License.