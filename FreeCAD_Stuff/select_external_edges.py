import FreeCAD
import FreeCADGui

# === Parameters ===
target_face_name = "Face1748"

# === Get selection ===
sel = FreeCADGui.Selection.getSelectionEx()
if not sel or not sel[0].HasSubObjects:
    raise Exception("Please select any edge or face.")

obj = sel[0].Object
shape = obj.Shape

# Get target face from name
face_index = int(target_face_name.replace("Face", "")) - 1
face = shape.Faces[face_index]

# Use OuterWire to get external boundary
outer_edges = face.OuterWire.Edges

# Select these edges in the GUI
FreeCADGui.Selection.clearSelection()
for edge in outer_edges:
    for i, e in enumerate(shape.Edges):
        if edge.isSame(e):
            FreeCADGui.Selection.addSelection(obj, f"Edge{i+1}")

print(f"Selected {len(outer_edges)} outer boundary edge(s) on {target_face_name}.")
