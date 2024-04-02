document.addEventListener('DOMContentLoaded', function() {
    var colorChoice = localStorage.getItem('colorChoice');
    if (colorChoice) {
        changeColor(colorChoice);
        var menu = document.getElementById('menuboard');
        menu.style.backgroundColor = colorChoice;
    }
});

function changeColor(newColor) {
    var submit = document.getElementById('submit');
    
    submit.classList.remove('blue-bg', 'red-bg', 'green-bg', 'yellow-bg');

    switch (newColor) {
        case '#3983C3':
            submit.classList.add('blue-bg');
            break;
        case '#CD2C2C':
            submit.classList.add('red-bg');
            break;
        case '#469C49':
            submit.classList.add('green-bg');
            break;
        case '#EBC842':
            submit.classList.add('yellow-bg');
            break;
        default:
            break;
    }

    localStorage.setItem('colorChoice', newColor);
}