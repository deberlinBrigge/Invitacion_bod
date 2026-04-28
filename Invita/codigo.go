function doGet() {
  return HtmlService.createTemplateFromFile('Index')
      .evaluate()
      .setTitle('Invitación: Lorenzo & Thalia')
      .setXFrameOptionsMode(HtmlService.XFrameOptionsMode.ALLOWALL)
      .addMetaTag('viewport', 'width=device-width, initial-scale=1');
}

// Esta función es por si luego quieres guardar confirmaciones en tu "Base_Gym"
function guardarAsistencia(datos) {
  const ss = SpreadsheetApp.getActiveSpreadsheet();
  const hoja = ss.getSheetByName("Usuario"); // Asegúrate de que el nombre coincida con tu Excel
  hoja.appendRow([new Date(), datos.nombre, "Invitado", datos.pases]);
  return "¡Registro completado con éxito!";
}