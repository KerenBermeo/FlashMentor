package models

type OpenQuestions struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	TopicID  uint   `gorm:"not null"`
	LevelID  uint   `gorm:"not null"`
	Question string `gorm:"type:varchar(300);not null"`
	Answer   string `gorm:"type:text;not null"`
	UserID   uint   `gorm:"not null"`
	Topic    Topics `gorm:"foreignKey:TopicID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Level    Level  `gorm:"foreignKey:LevelID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User     User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}
