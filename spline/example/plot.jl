using Winston

Data = readcsv("Data/test.csv")

X = Data[:,1]
Y = Data[:,2]

F = FramedPlot(
    title="Spline",
    xlabel="x",
    ylabel="y"
)

x = collect(0.:1.:40.);
y = sin.(x);

P = Points(x,y,size=.1);
C = Curve(X, Y)

setattr(P, "label", "points")
setattr(C, "label", "Spline")
lgnd = Legend(.05, .1, [P, C])

add(F, P, C, lgnd)
savefig(F, "Fig/spline.svg", (1000,600))
run(`inkscape -z Fig/spline.svg -e Fig/spline.png -d 300 --export-background=WHITE`)
run(`rm Fig/spline.svg`)