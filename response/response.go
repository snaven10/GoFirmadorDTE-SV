package response

// ResponseBody representa la estructura de respuesta
type ResponseBody struct {
	Status string      `json:"status"`
	Body   interface{} `json:"body"`
}

// Mensaje es una estructura que contiene los métodos para crear respuestas
type Mensaje struct{}

// NewMensaje crea una instancia de Mensaje
func NewMensaje() *Mensaje {
	return &Mensaje{}
}

// OK crea una respuesta "ok" con el cuerpo especificado
func (m *Mensaje) OK(body interface{}) ResponseBody {
	return ResponseBody{Status: "ok", Body: body}
}

// Error crea una respuesta de error con el código y mensaje especificados
func (m *Mensaje) Error(codigo string, mensaje interface{}) ResponseBody {
	return ResponseBody{Status: "error", Body: BodyMensaje{Codigo: codigo, Mensaje: mensaje}}
}

// BodyMensaje representa la estructura del cuerpo del mensaje de error
type BodyMensaje struct {
	Codigo  string      `json:"codigo"`
	Mensaje interface{} `json:"mensaje"`
}
