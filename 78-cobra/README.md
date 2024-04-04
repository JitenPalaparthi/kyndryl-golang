# How to develop CLI tools

1. directly use os.Args : Write lot of code to parse values , commands, flags etc..
2. flag package: it is used to parse flags but does not completely support commands and sub commands.
3. user cobra: It is a third party package. It eases creating CLI based applicaions with the support of Commands/Sub-Commands/Flags and also help doc.