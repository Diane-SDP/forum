document.addEventListener('DOMContentLoaded', function() {
  const mediaInput = document.getElementById('mediaInput');
  const uploadedImage = document.getElementById('uploadedImage');
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

  // Gestionnaire d'événements pour supprimer l'image
  removeImageBtn.addEventListener('click', function(e) {
    e.preventDefault();
      uploadedImage.src = ''; // Effacer le contenu de l'image
      uploadedImage.style.display = 'none'; // Cacher l'image
      removeImageBtn.style.display = 'none'; // Cacher le bouton "Supprimer"
      mediaInput.value = ''; // Effacer la valeur de l'input de type file
  });
});
