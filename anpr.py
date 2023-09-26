from flask import Flask, request
import os, sys
import datetime
# import mysql.connector
import xmlrpc.client
import json

app = Flask(__name__)
# connection = mysql.connector.connect(host='192.168.1.68',user='pcs',password='123456',database='notify')

@app.route('/api/anprNotify', methods=['POST'])
def notify():
    content_type = request.headers.get('Content-Type')
    if (content_type == 'application/json'):
        jsondata = request.json
        # writeToLog(jsondata)
        writeToRaw(jsondata)
        
        waktu = jsondata['time']
        plate = jsondata['plate']
        writeToLog("waktu : " + str(waktu) + " | plate : " + str(plate))
        s = xmlrpc.client.ServerProxy('http://localhost:20008/RPC2')
        listMethods = s.system.listMethods()
        # print(arraystruct)
        result = s.ANPR.Data(waktu, plate)
        # return result
        # return listMethods
        return { "rcode" : 0, "status" : "Ok", "waktu" : str(waktu), "plate" : str(plate), "result" : str(result) }
    else:
        return { "rcode" : 1, "status" : "Invalid Format"}
    # return "Hello World!"

def writeToRaw(raw):
    os.environ['TZ'] = 'Asia/Jakarta'
    # today = datetime.datetime.now().strftime("%Y-%m-%d")
    pathraw = '/home/module/tct/anprNotify/raw'
    mylog = open(pathraw,'w+')
    mylog.write(str(raw))
    mylog.write("\n\n")

def writeToLog(raw):
    os.environ['TZ'] = 'Asia/Jakarta'
    today = datetime.datetime.now().strftime("%Y-%m-%d")
    pathraw = '/home/module/tct/anprNotify/Logs/' + '[' + today + '].log'
    mylog = open(pathraw,'a+')
    mylog.write('[' + str(datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S")) + '] ')
    mylog.write(str(raw))
    mylog.write("\n")

def writeToDB(sql, connection):
    # writeToLog("user mysql => " + USER)
    mycursor = connection.cursor()
    mycursor.execute(sql)
    connection.commit()
    mycursor.close()

if __name__ == '__main__':
    app.run(host='0.0.0.0', port='22222', debug=False)
