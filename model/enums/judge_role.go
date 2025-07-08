package enums

type JudgeRole string

const (
	JudgeRoleANN  JudgeRole = "ANN"  //	Announcers or speaker
	JudgeRoleAREF JudgeRole = "AREF" //	Assistant referee
	JudgeRoleCAR  JudgeRole = "CAR"  //	Callroom
	JudgeRoleCARS JudgeRole = "CARS" //	Callroom supervisor
	JudgeRoleCFIN JudgeRole = "CFIN" //	Chief finish judge
	JudgeRoleCIOT JudgeRole = "CIOT" //	Chief inspectors of turns
	JudgeRoleCOC  JudgeRole = "COC"  //	Clerks of course
	JudgeRoleCR   JudgeRole = "CR"   // Control room (computerroom)
	JudgeRoleCREC JudgeRole = "CREC" //	Chief recorder
	JudgeRoleCRS  JudgeRole = "CRS"  //	Control room supervisor
	JudgeRoleCTIK JudgeRole = "CTIK" //	Chief timekeeper
	JudgeRoleFIN  JudgeRole = "FIN"  //	Finish judge
	JudgeRoleFSR  JudgeRole = "FSR"  //	False start rope personnel.
	JudgeRoleIOT  JudgeRole = "IOT"  //	Inspector of turns
	JudgeRoleJOS  JudgeRole = "JOS"  //	Judge of strokes
	JudgeRoleMDR  JudgeRole = "MDR"  //	Meet director
	JudgeRoleMED  JudgeRole = "MED"  //	Medical service
	JudgeRoleOTH  JudgeRole = "OTH"  //	other or unknown
	JudgeRoleREC  JudgeRole = "REC"  //	Recorder
	JudgeRoleREF  JudgeRole = "REF"  //	Referee
	JudgeRoleSFO  JudgeRole = "SFO"  //	Safety officer
	JudgeRoleSTA  JudgeRole = "STA"  //	Starter
	JudgeRoleTDG  JudgeRole = "TDG"  //	Technical delegate
	JudgeRoleTIK  JudgeRole = "TIK"  //	Timekeeper
	JudgeRoleWAS  JudgeRole = "WAS"  //	Warmup supervisor
)
