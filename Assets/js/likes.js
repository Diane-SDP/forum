function Like(e) {
    url = "/"
    displayLikes = document.getElementById('nbLikes' + e.value)
    displayIMG = document.getElementById('Likes' + e.value).getElementsByTagName('img')[0]
    console.log(displayIMG.src)
    if (displayIMG.src == "http://localhost:8080/Assets/img/likes.png") {
        var data = {
            action:"dislike",
            id : e.value,
        }
        value = parseInt(displayLikes.innerText)
        displayLikes.innerText = value - 1
        displayIMG.src = "http://localhost:8080/Assets/img/trending-topic%201.png"
    } else {
        var data = {
            action:"like",
            id : e.value,
        }
        value = parseInt(displayLikes.innerText)
        displayLikes.innerText = value + 1
        displayIMG.src = "http://localhost:8080/Assets/img/likes.png"
    }
    fetch(url, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    })
    .catch(error => {
        console.error('Erreur lors de l\'envoi de la requête:', error);
    });
}