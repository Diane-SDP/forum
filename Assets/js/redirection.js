function redirectionChronometree(url, delai) {
    setTimeout(function() {
      window.location.href = url;
    }, delai);
  }

  // Rediriger vers une autre page après 2 secondes (2000 millisecondes)
  redirectionChronometree('/', 2000);