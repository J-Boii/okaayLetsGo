count = 0
def main():
    global count
    if count == 995:
        return
    else:    
        count += 1
        print(count)
        main()

main()