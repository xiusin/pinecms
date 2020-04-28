// JavaScript Document

function hot7()
{
    document.getElementById("con1").style.display = "";
    document.getElementById("con2").style.display = "none";
    document.getElementById("con3").style.display = "none";
    document.getElementById("con4").style.display = "none";
    document.getElementById("hot7").className = "current";
    document.getElementById("comm7").className = "";
    document.getElementById("month").className = "";
    document.getElementById("must").className = "";
}

function comm7() {
    document.getElementById("con1").style.display = "none";
    document.getElementById("con2").style.display = "";
    document.getElementById("con3").style.display = "none";
    document.getElementById("con4").style.display = "none";
    document.getElementById("hot7").className = "";
    document.getElementById("comm7").className = "current";
    document.getElementById("month").className = "";
    document.getElementById("must").className = "";
}

function month() {
    document.getElementById("con1").style.display = "none";
    document.getElementById("con2").style.display = "none";
    document.getElementById("con3").style.display = "";
    document.getElementById("con4").style.display = "none";
    document.getElementById("hot7").className = "";
    document.getElementById("comm7").className = "";
    document.getElementById("month").className = "current";
    document.getElementById("must").className = "";
}

function must() {
    document.getElementById("con1").style.display = "none";
    document.getElementById("con2").style.display = "none";
    document.getElementById("con3").style.display = "none";
    document.getElementById("con4").style.display = "";
    document.getElementById("hot7").className = "";
    document.getElementById("comm7").className = "";
    document.getElementById("month").className = "";
    document.getElementById("must").className = "current";
}