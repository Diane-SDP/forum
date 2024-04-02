document.addEventListener('DOMContentLoaded', function() {
    var colorChoice = localStorage.getItem('colorChoice');
    if (colorChoice) {
        var header = document.getElementById('header');
        header.style.backgroundColor = colorChoice;
        var menu = document.getElementById('menuboard');
        menu.style.backgroundColor = colorChoice;
    }
});

function changeColor(newColor) {
    var header = document.getElementById('header');
    header.style.backgroundColor = newColor;
    var menu = document.getElementById('menuboard');
    menu.style.backgroundColor = newColor;

    localStorage.setItem('colorChoice', newColor);
}

