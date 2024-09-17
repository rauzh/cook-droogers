#!/opt/homebrew/bin/gnuplot -persist

set style data histograms
set style histogram clustered
set style fill solid
set boxwidth 0.9 relative
set xtics nomirror
set ytics nomirror
set grid ytics
set style line 1 linecolor rgb 'skyblue'
set style line 1 linecolor rgb 'salmon'

set key top right

set xlabel 'Число потоков, шт.'
set ylabel 'Время выполнения, мс'

plot 'data_restore' using 2:xtic(1) title 'pg\_restore'