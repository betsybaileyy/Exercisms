
def two_fer(name):

    if name is not None:
        txt = "One for {}, one for me".format(name)
        # print("One for {}, one for me.".format(name))
        print(txt)

    else:
        txt_2 = "One for you, one for me"
        # print("One for you, one for me.")
        print(txt_2)

two_fer()
