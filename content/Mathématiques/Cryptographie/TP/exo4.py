from exo2 import bezout


# DOIT ÊTRE PREMIER SINON NE MARCHE PAS
dict_cod={'a':1,'b':2,'c':3,'d':4,'e':5,'f':6,'g':7,'h':8,'i':9,'j':10,'k':11,'l':12,'m':13,'n':14,'o':15,'p':16,'q':17,\
                   'r':18,'s':19,'t':20,'u':21,'v':22,'w':23,'x':24,'y':25,'z':26,' ':27,',':28,'.':29,'?':30, ':':0}

encoded_dict = dict_cod

def encode(text, a, b): # chiff_afine
    encodedVal = []
    encodedText = ""
    
    for i in text :
        encodedVal.append((a*dict_cod[i]+b))
    for i in encodedVal :
        for key, value in dict_cod.items() :
            # print(key, value, i, i%31)
            if value == i%len(dict_cod) :
                encodedText+=key

    return encodedText

def decode(mess, a, b):
    encodedVal = []
    decodedText = ""
    
    euclide_etendu = bezout(a, len(dict_cod))
    a = euclide_etendu[1] # coef de bezout associé

    for char in mess :
        # m = a**(-1) * (c-b) [32] # a-> coef de bezout
        charCode = a*(dict_cod[char]-b)%len(dict_cod)

        wasDecode = False
        for key, value in dict_cod.items() :
            # print(key, value, i, i%31)
            if value == charCode%len(dict_cod) :
                decodedText+=key
                wasDecode = True
        if not(wasDecode) :
            print("failed to decode", char, ":", charCode)

    return decodedText

def encodeDecodeTest(msg, a, b):
    code = encode(msg, a, b)
    nmsg = decode(code, a, b)
    print(msg == nmsg, ":", msg, "->", code, "->", nmsg)


if __name__=="__main__":
    for i in range(0,100):
        for j in range(0,100):
            encodeDecodeTest("hello wolrd", i, j)
