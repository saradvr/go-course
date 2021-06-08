package testPackage

func (metodo *Curso) GetTitulo() string { // si el nombre está en mayúscula es público, si empieza con minúscula, es privado
	return metodo.Titulo
}
