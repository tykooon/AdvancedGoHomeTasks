// ****************************  TODO App v.0.1.0 *********************************
// 
//    This project includes two services (todoapp & todoweb) with a shared resource 
// file: todolist.json.
//    In order to use both services in parallel it's highly recommended to 
// compile both parts of the project into a common directory (e.g. the root directory
// of the project).
//    In the root directory you can find a batch file, that will simplify the compilation
// process: compiler.bat
//    This file is created for your convenience and runs two compilation commands
// from the root:
// 		>>go build ./cmd/todoapp	
// 		>>go build ./cmd/todoweb
//  
// This way you can be sure that both services will work with the same shared directory.
//
// **********************************************************************************