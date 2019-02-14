# tex-pack-helper
A selection of tools written in Go for helping with the generation and post-processing of neural network-enhanced video game textures.

## Purpose
Using neural networks such as waifu2x is becomming a popular way to breath new life into older games by upscaling those textures so that they look much better at higher resolutions. This process isn't perfect, however, and could benefit from additional automation. 

### Issues With Using Neural Networks and Game Textures
<ul>
  <li>Games use a huge number of textures. Even games from the GameCube, PS2, and XBox era have thousands of textures for a single game. To run those through waifu2x, you have to create a list of all of the textures to convert, which could take an obscenely long time if done by hand.</li>
  <li>Neural networks are far from perfect and will produce undesireable results at times. </li>
  <ul>
    <li>Mask textures may be output without any transparency.</li>
    <li>Textures using a single color plus full transparency may be output as only the solid color.</li>
    <li>The neural network may make undesireable assumptions when upscaling.</li>
  </ul>
</ul>

### Solutions Provided by tex-pack-helper
* It can automatically make a list of the files to use in the neural net.
* Output files are analyzed and reclassified based on common output issues.

## Important Notes
* tex-pack-helper is still in extremely early development.
* User-friendliness will improve as the project progresses.

## In Progress
* Instructions on how to install the software
* Instructions on how to contribute
* Image analysis feature
