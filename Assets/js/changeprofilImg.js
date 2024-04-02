document.addEventListener('DOMContentLoaded', function() {
    const mediaInput = document.getElementById('img');
    const uploadedImage = document.getElementById('uploadedImage');
  
    mediaInput.addEventListener('change', function(event) {
      const file = event.target.files[0]; 
      const reader = new FileReader();
  
      reader.onload = function(e) {
        uploadedImage.style.backgroundImage = `url(${e.target.result})`;
        uploadedImage.style.backgroundSize = 'cover'; 
        uploadedImage.style.backgroundPosition = 'center'; 
        uploadedImage.style.backgroundColor = 'transparent';
        uploadedImage.setAttribute('src', '../Assets/img/none.png'); 
      };
  
      if (file) {
        reader.readAsDataURL(file); 
      }
    });
  });