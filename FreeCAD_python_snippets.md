### open file in gui
```
from PySide import QtGui
name=QtGui.QFileDialog.getOpenFileName()[0]
if len(name) > 0:
    txtFile = open(name,"r")
    content = txtFile.readlines()
```
