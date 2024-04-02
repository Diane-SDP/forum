const imgs = document.querySelectorAll('.content-post img');
const fullPage = document.querySelector('#fullpage');

imgs.forEach(img => {
  img.addEventListener('click', function() {
      fullPage.style.backgroundImage = 'url(' + img.src + ')';
      fullPage.style.display = 'block';
  });
});

document.addEventListener('DOMContentLoaded', function () {
    const postImages = document.querySelectorAll('#postImage');

    postImages.forEach(postImage => {
        if (!postImage || !postImage.complete || postImage.naturalWidth === 0) {
            postImage.style.display = 'none';
        } else {
            postImage.style.display = 'block';
        }
    });
});