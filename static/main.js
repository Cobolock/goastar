var canvas = document.getElementById("canvas")
var cnv = canvas.getContext("2d")
var drawing = false
var current = 0
var obstacles = []

function redraw(){
    //Clear all
    cnv.clearRect(0, 0, canvas.width, canvas.height)
    //Outer borders
    cnv.strokeRect(0, 0, canvas.width, canvas.height)
    //Startng point
    cnv.beginPath()
    cnv.arc(50, 550, 5, 0, 2 * Math.PI, false)
    cnv.fillStyle = 'navy'
    cnv.fill()
    cnv.closePath()
    //Ending point
    cnv.beginPath()
    cnv.arc(850, 50, 5, 0, 2 * Math.PI, false)
    cnv.fillStyle = 'green'
    cnv.fill()
    cnv.closePath()
    //Redraw obstacles
    obstacles.forEach(function(ob, i, obstacles){
        cnv.beginPath()
        cnv.moveTo(ob[0].x, ob[0].y)
        ob.forEach(function(point, j, ob){
            cnv.lineTo(point.x, point.y)
        })
        cnv.stroke()
        cnv.fill()
        cnv.closePath()
    })
}

canvas.onmousemove = function(e){
    if (drawing) {
        obstacles[current].push({"x" : e.layerX, "y": e.layerY})
        redraw()
    }
}

canvas.onmousedown = function() {
    drawing = true
    current = obstacles.push([]) - 1
}

canvas.onmouseup = function(){
    drawing = false
    if (obstacles[current].length < 2) {
        obstacles.pop()
    }
    current = null
}

function send() {
    var xmlhttp = new XMLHttpRequest()
    xmlhttp.open('POST', './js')
    xmlhttp.setRequestHeader('Content-Type', 'application/json;charset=UTF-8')
    xmlhttp.send(JSON.stringify(obstacles))
}

redraw()
