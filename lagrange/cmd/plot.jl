using Winston

Data = readcsv("Data/Lagrange.csv")

X = Data[:,1];
Y = Data[:,2];

Px = [1, 2, 3, 4]
Py = [1, 4, 9, 16]

p = FramedPlot(
    title = "Lagrange Interpolation",
    xlabel="X",
    ylabel="Y",
    xrange=(0,5),
    yrange=(0,25)
)
C = Curve(X, Y)
P = Points(Px, Py, symbolkind="plus")
setattr(C, "label", "x\^2")
setattr(P, "label", "Point")
lgnd = Legend(.1, .9, [C, P])
add(p, C, lgnd)
add(p, P)
savefig(p, "Fig/lagrange.svg", (1000,600))

run(`inkscape -z Fig/lagrange.svg -e Fig/lagrange.png -d 300 --export-background=WHITE`)