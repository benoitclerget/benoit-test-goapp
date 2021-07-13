
// first page loading: show list of employess
document.addEventListener("DOMContentLoaded", async function() {
  if (window.fetch) {
    try{
      let response = await fetch('/api/system_infos')
      if(response.ok) {
        response.text().then(function(myText) {
            // set values into the dom
            let app_div_webapp_infos_elem = document.getElementById('app_webapp_infos');
            app_div_webapp_infos_elem.innerHTML = `<span class="app--type-mono-scale_4"><a href="https://github.com/benoitclerget/benoit-test-goapp" target="_blank">Git repository</a></span>`
            let app_div_system_infos_elem = document.getElementById('app_system_infos');
            app_div_system_infos_elem.innerHTML = `<span class="app--type-mono-scale_4">${myText}</span>`
        });
      } else {
        console.error('response error: ' + response.ok);
      }
  
    } catch(error) {
      console.error('Error with fetch operation: ' + error);
    }
  } else {
    console.error('fetch not supported');
  }

})
