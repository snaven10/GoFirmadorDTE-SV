# Firmador para Factura Electrónica en El Salvador basado en la solución gratuita que ofrece MH para firmar DTE

Herramienta diseñada específicamente para firmar documentos de factura electrónica en El Salvador, siguiendo los estándares y requisitos establecidos por el Ministerio de Hacienda (MH). Basada en la solución gratuita que ofrece MH para firmar DTE y construida con Go para garantizar rapidez y eficiencia.

## Características Principales

- **Diseñado para El Salvador**: Cumple con todas las regulaciones y requisitos específicos de factura electrónica del Ministerio de Hacienda de El Salvador.
- **Implementación en Go**: Aprovecha la rapidez y el rendimiento de Go para procesar y firmar facturas electrónicas de manera eficiente.
- **Uso de `go-jose/v3`**: Utiliza la librería [go-jose/v3](https://github.com/go-jose/go-jose/) para garantizar firmas seguras y compatibles.
- **Verificación de RSA**: Asegura que solo se usen claves RSA para la firma de facturas electrónicas.
- **Lectura de Certificados**: Capacidad para leer y procesar certificados usados en la firma de facturas electrónicas.
- **Control de Errores**: Gestión avanzada de errores para garantizar un manejo adecuado y comunicación clara al usuario.
- **Desarrollo con DevContainer**: Este proyecto utiliza DevContainer para proporcionar un entorno de desarrollo consistente y aislado.

## Uso del Firmador

Para instrucciones detalladas sobre cómo usar el firmador para facturas electrónicas en El Salvador, visite el sitio oficial Facturación Electronica: [informacion tecnica y funcional](https://factura.gob.sv/informacion-tecnica-y-funcional/).

## Instalación

```shell
# Clone el repositorio
git clone https://github.com/snaven10/GoFirmadorDTE-SV

# Entre al directorio del proyecto
cd GoFirmadorDTE-SV

# Instale las dependencias
go get
```

## Contribuciones

Las contribuciones son bienvenidas. Si encuentra algún error o desea agregar alguna mejora específica para El Salvador, no dude en hacer un pull request o abrir un issue.
