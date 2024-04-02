document.addEventListener('DOMContentLoaded', function() {
    const mediaInput = document.getElementById('mediaInput');
    const uploadedImage = document.getElementById('postImage');
    const removeImageBtn = document.getElementById('removeImage');
  
    mediaInput.addEventListener('change', function(event) {
        const file = event.target.files[0];
        const reader = new FileReader();
  
        reader.onload = function(e) {
            uploadedImage.src = e.target.result;
            uploadedImage.style.display = 'block';
            removeImageBtn.style.display = 'flex'; // Afficher le bouton "Supprimer"
        };
        reader.readAsDataURL(file);
    });
  
    if (uploadedImage.src !== '.' && uploadedImage.src.includes('.')) {
        removeImageBtn.style.display = 'flex';
    } else {
        removeImageBtn.style.display = 'none';
    }
    
    // supprimer l'image
    removeImageBtn.addEventListener('click', function(e) {
      e.preventDefault();
        uploadedImage.src = ''; // Efface le contenu de l'image
        uploadedImage.style.display = 'none'; // Cache l'image
        removeImageBtn.style.display = 'none'; // Cache la croix
        mediaInput.value = ''; // Effacer la valeur de l'input de type file
    });
  });
  