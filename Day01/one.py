f = open('input.txt')
inputlist = f.read().splitlines()
f.close()
inccount = 0
slidingwindowinc = 0
slidingwindow = []
for n in range(0, len(inputlist) - 1):
    if  n < len(inputlist) - 2:
        slidingwindow.append(int(inputlist[n]) + int(inputlist[n+1]) + int(inputlist[n+2]))
    if int(inputlist[n]) < int(inputlist[n+1]):
        inccount += 1
for x in range(0, len(slidingwindow) - 1):
    if slidingwindow[x] < slidingwindow[x+1]:
        slidingwindowinc += 1

print(inccount)
print(slidingwindowinc)
