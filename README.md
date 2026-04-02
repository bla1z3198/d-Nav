# dnav

Program for creating and calculating trajectories of aircraft. 
which contains curves and lines.

# v3.1
Last version (v3.1) have commands and enhanced system of upload to a file
* init [filename] - init json with filename
* nav - calculate trajectory from current json
* upload [filename] - upload flight plan into file
* exit - exit from dnav

## You can check work of 'upload' command in upload.txt
# Examples
Example of routes stores in data.json and data2.json

# Doc

1 - Core of trajectory engine.
    
    1.1 - Structure of waypoint.
        Every waypoint have:
            - X coordinate
            - Y coordinate
            - Z coordinate
            - ID
            - R radius of circle
            around waypoint.
    
    1.2 - Navigation func

        1.2.1 - Variables
        - prev_sin float64 - Sinus of angle
        between "x" axis and line which connect
        A and B point (previous)
        - prev_cos float64 - Cosinus of same angle
        (previous)
        - result float64 - Result length of way
        - tangent float64 - Length of line between
        "start" point and first point in json
        - tangent_line float64 - Length of line
        between two "tangent" points
        - sin float64 - Same to prev_sin, but
        this is current value of sinus
        - cos float64 - Same to prev_cos, but
        this is current value of cosinus
        - angle0 float64 - Angle between x axis and
        first line
        - x float64 - X coordinate of start tangent
        point
        - y float64 - Y coordinate of start tangent
        point
        - arg float64 - Argument for arccos of 
        sector angle
        - gamma float64 - Angle of circle sector
        - sector float54 - Length of circle sector

        1.2.2 - Geometry... 
	(there will be explanation in next commit)
