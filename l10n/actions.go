package l10n

import "github.com/wt-tools/wtscope/action"

type translatedAction struct {
	Code action.Code
	List []translation
}

var actionTexts = []translatedAction{
	{
		action.Unknown,
		[]translation{
			{En, "unknown action"},
			{Ru, "неизвестное действие"},
		},
	},
	{
		action.Destroyed,
		[]translation{
			{En, "destroyed"},
			{Ru, "уничтожил"},
		},
	},
	{
		action.Wrecked,
		[]translation{
			{En, "has been wrecked"},
			{Ru, "выведен из строя"},
		},
	},
	{
		action.Achieved,
		[]translation{
			{En, "has achieved"},
			{Ru, "получил"},
		},
	},

	// {En, action.Unknown, "has joined event"},
	{
		action.Afire,
		[]translation{
			{En, "set afire"},
			{Ru, "поджег"},
		},
	},
	{
		action.Crashed,
		[]translation{
			{En, "has crashed."},
			{Ru, "разбился"},
		},
	},
	{
		action.Connected,
		[]translation{
			{En, "connected"},
			{Ru, "присоединился"},
		},
	},
	{
		action.SoftLanding,
		[]translation{
			{En, "performed a soft landing"},
		},
	},
	{
		action.FinalBlow,
		[]translation{
			{En, "has delivered the final blow!"},
			{Ru, "нанёс последний удар!"},
		},
	},
	{
		action.Damaged,
		[]translation{
			{En, "damaged"},
			{Ru, "подбил"},
		},
	},
	{
		action.ShotDown,
		[]translation{
			{En, "shot down"},
			{Ru, "сбил"},
		},
	},
	{
		action.NetworkDisconnect,
		[]translation{{Auto, "td! kd?NET_PLAYER_DISCONNECT_FROM_GAME"}},
	},
	{
		action.Disconnected,
		[]translation{
			{En, "has disconnected from the game"},
			{Ru, "потерял связь"},
		},
	},
}
