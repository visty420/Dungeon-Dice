if command -v gnome-terminal &> /dev/null
then
    gnome-terminal -- bash -c "./dungeondice; exec bash"
elif command -v xterm &> /dev/null
then 
    xterm -fa 'Monospace' -fs 14 -bg black -fg white -fullscreen -e "./dungeondice"
else
    echo "No supported terminal found. Please install gnome-terminal or xterm."
fi