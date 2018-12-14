package utils

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
)

// Path to config
var (
	home, _     = homedir.Dir()
	BroConfPath = filepath.Join(home, ".brocli")
)

// Constants
const (
	NLn = "\n      " // New line with indentation

	UserInitH = `#pragma once

// Include Game levels here as
// #include "game_name/engine_files/level.level_name.h"
`
	UserInitCpp = `#include "user_init.h"
#include "../RubeusCore.h"

// Initialise levels here as
// SampleLevel * sample_level = new SampleLevel("sample_level");

// Add startup level as
// std::string startupLevel = "sample_level";

// Initialise game objects as
// SampleObject * sample_object = new SampleObject(
// 	   "sample_object",
// 	   "sample_level",
// 	   6.0f, 3.0f,
// 	   3.0f, 3.0f,
// 	   "Assets/debug.png",
// 	   false
// );
`
	LevelH = `#pragma once

#include "../RubeusCore.h"

class %s : public Rubeus::RLevel
{
protected:
	// User defined members


	// Don't look down
public:
	%s(std::string name)
		: RLevel(name)
	{
	}

	~%s()
	{
	}

	void begin() override;
	void onEnd() override;
};
`
	LevelCpp = `#include "%s.h"

void %s::begin()
{
}

void %s::onEnd()
{
}
`
	ObjectH = `#pragma once

#include "../RubeusCore.h"

class %s : public Rubeus::RGameObject
{
	// User members



	// Do not look below
public:

	%s(std::string name, std::string levelName, float x, float y, float deltaX, float deltaY, const char * imageFilePath, bool enablePhysics = false, const Rubeus::Awerere::EColliderType & type = Rubeus::Awerere::EColliderType::NO_COLLIDER, Rubeus::Awerere::ACollider * collider = NULL, bool generatesHit = false, const Rubeus::Awerere::APhysicsMaterial & physicsMat = Rubeus::Awerere::APhysicsMaterial())
		: RGameObject(name, levelName, x, y, deltaX, deltaY, imageFilePath, enablePhysics, type, collider, generatesHit, physicsMat)
	{
	}

	// TODO: Add default paramters
	%s(std::string name, std::string levelName, float x, float y, float deltaX, float deltaY, float & r, float & g, float & b, bool enablePhysics = false, const Rubeus::Awerere::APhysicsMaterial & material = Rubeus::Awerere::APhysicsMaterial(), const Rubeus::Awerere::EColliderType & type = Rubeus::Awerere::EColliderType::NO_COLLIDER, Rubeus::Awerere::ACollider * collider = NULL, bool generatesHit = false)
		: RGameObject(name, levelName, x, y, deltaX, deltaY, r, g, b, enablePhysics, material, type, collider, generatesHit)
	{
	}

	void begin() override;
	void onHit(RGameObject * hammer, RGameObject * nail, const Rubeus::Awerere::ACollideData & collisionData) override;
	void onMessage(Rubeus::Message * msg) override;
	void tick() override;
};
`

	ObjectCpp = `#include "%s.h"

void %s::begin()
{
}

void %s::onHit(RGameObject * hammer, RGameObject * nail, const Rubeus::Awerere::ACollideData & collisionData)
{
}

void %s::tick()
{
}

void %s::onMessage(Rubeus::Message * msg)
{
}
`
)
