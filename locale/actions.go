package locale

import "github.com/wt-tools/hq/action"

var actionTexts = []Translation{
	{En, action.Unknown, "unknown"},
	{Ru, action.Unknown, "неизвестное действие"},

	{En, action.Destroyed, "destroyed"},
	{Ru, action.Destroyed, "уничтожил"},

	{En, action.Wrecked, "has been wrecked"},
	{Ru, action.Wrecked, "выведен из строя"},

	{En, action.Achieved, "has achieved"},
	{Ru, action.Achieved, "получил"},

	{En, action.Unknown, "has joined event"}, // XXX

	{En, action.Afire, "set afire"},
	{Ru, action.Afire, "поджег"},

	{En, action.Connected, "connected"},
	{Ru, action.Connected, "присоединился"},

	{En, action.SoftLanding, "performed a soft landing"},

	{En, action.FinalBlow, "has delivered the final blow!"},
	{Ru, action.FinalBlow, "нанёс последний удар!"},

	{En, action.Damaged, "damaged"},

	{En, action.ShotDown, "shot down"},
	{Ru, action.ShotDown, "сбил"},

	{Ru, action.Unknown, "подбил"}, // XXX

	{Ru, action.Unknown, "разбился"}, // XXX

	{En, action.Disconnected, "td! kd?NET_PLAYER_DISCONNECT_FROM_GAME"},

	{Ru, action.LostConnect, "потерял связь"}, // XXX
}
