package locale

import "github.com/wt-tools/adjutant/damage"

var damageTexts = []Translation{
	{En, damage.Unknown, "unknown"},
	{Ru, damage.Unknown, "неизвестное действие"},

	{En, damage.Destroyed, "destroyed"},
	{Ru, damage.Destroyed, "уничтожил"},

	{En, damage.Wrecked, "has been wrecked"},
	{Ru, damage.Wrecked, "выведен из строя"},

	{En, damage.Got, "?"},
	{Ru, damage.Got, "получил"},

	{En, damage.Fired, "fired"},
	{Ru, damage.Fired, "поджег"},

	{En, damage.Connected, "connected"},
	{Ru, damage.Connected, "присоединился"},
}
