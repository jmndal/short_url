function shortenURL() {
  var longURL = document.getElementById("urlInput").value;
  console.log("longURL: " + longURL);

  fetch("/shorten/", {
    method: "POST",
    headers: {
      "Content-Type": "application/x-www-form-urlencoded",
    },
    body: "url=" + encodeURIComponent(longURL),
  })
    .then((response) => {
      if (!response.ok) {
        throw new Error("Failed to shorten URL");
      }
      return response.text();
    })
    .then((data) => {
      // Display the shortened URL as a link
      console.log("DATA: " + data);
      document.getElementById("shortenedURL").innerHTML = "Shortened URL: <input id='shortenedLink' value='" + data + "' readonly></input>";
    });
}

function redirectURL(shortenedURL) {
  fetch("/redirect/" + encodeURIComponent(shortenedURL))
    .then((response) => {
      if (!response.ok) {
        throw new Error("Failed to redirect");
      }
      return response.text();
    })
    .then((originalURL) => {
      console.log("original: " + originalURL);
      // Redirect to the original URL
      window.location.href = originalURL;
    })
    .catch((error) => {
      console.error("Redirect error:", error);
    });
}