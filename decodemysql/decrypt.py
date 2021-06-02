#!/usr/local/bin/python2.7
#coding:utf8

import sys,base64,pyDes
def __encrypt(data):

    try:
        k = pyDes.des('whGcZugH', pyDes.ECB, '\0'*8, pad=None, padmode=pyDes.PAD_PKCS5)
        d = base64.b64decode(data)
        ret = k.decrypt(d)
        return ret
    except Exception, e:
        return 1, e

s=open(sys.argv[1]).readline()
print __encrypt(s.strip())