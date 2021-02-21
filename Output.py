import unreal

asset_path = "/Game/Materials/"
root = unreal.AssetToolsHelpers.get_asset_tools().create_asset("PannerFunction", asset_path, unreal.Material, unreal.MaterialFactoryNew())
root.set_editor_property("d3d11_tessellation_mode", unreal.MaterialTessellationMode.MTM_PN_TRIANGLES)
matref.set_editor_property("automatically_set_usage_in_editor", False)
matref.set_editor_property("used_with_instanced_static_meshes", False)
matref.set_editor_property("used_with_particle_sprites", False)
matref.set_editor_property("used_with_niagara_mesh_particles", True)
matref.set_editor_property("used_with_water", False)
matref.set_editor_property("used_with_niagara_ribbons", True)
matref.set_editor_property("used_with_geometry_collections", False)
matref.set_editor_property("used_with_clothing", False)
matref.set_editor_property("used_with_hair_strands", False)
matref.set_editor_property("used_with_static_lighting", False)
matref.set_editor_property("used_with_morph_targets", False)
matref.set_editor_property("used_with_spline_meshes", False)
matref.set_editor_property("used_with_beam_trails", False)
matref.set_editor_property("used_with_mesh_particles", False)
matref.set_editor_property("used_with_editor_compositing", False)
matref.set_editor_property("used_with_skeletal_mesh", False)
matref.set_editor_property("used_with_niagara_sprites", True)
matref.set_editor_property("used_with_geometry_cache", False)

# Attach custom expression node
OTE4ZTdjMjg = unreal.MaterialEditingLibrary.create_material_expression(root, unreal.MaterialExpressionCustom, -400, 0)
OTE4ZTdjMjg.set_editor_property("description", "PannerFunction")
OTE4ZTdjMjg.set_editor_property("desc", "Similar to panner node. Flows in a direction.")
OTE4ZTdjMjg.set_editor_property("output_type", unreal.CustomMaterialOutputType.CMOT_FLOAT1)
OTE4ZTdjMjg.set_editor_property("code", '#include "/Panner.usf"\nreturn 0;')

# List of pins to attach
NTY4NjZlYmM = unreal.CustomInput()
NTY4NjZlYmM.set_editor_property("input_name", "size")
OWE1ZjRjZjk = unreal.MaterialEditingLibrary.create_material_expression(root, unreal.MaterialExpressionScalarParameter, -600, 0)
OWE1ZjRjZjk.set_editor_property("desc", "No description provided")
OWE1ZjRjZjk.set_editor_property("default_value", 10.000000)
OWE1ZjRjZjk.set_editor_property("parameter_name", "size")
MjIwMWRjMDk = unreal.CustomInput()
MjIwMWRjMDk.set_editor_property("input_name", "direction")
MzM0YTc5MGM = unreal.MaterialEditingLibrary.create_material_expression(root, unreal.MaterialExpressionScalarParameter, -600, 0)
MzM0YTc5MGM.set_editor_property("desc", "No description provided")
MzM0YTc5MGM.set_editor_property("default_value", 10.000000)
MzM0YTc5MGM.set_editor_property("parameter_name", "direction")
NjQzYjMwMWI = unreal.CustomInput()
NjQzYjMwMWI.set_editor_property("input_name", "offset")
OTlkYTczYWM = unreal.MaterialEditingLibrary.create_material_expression(root, unreal.MaterialExpressionScalarParameter, -600, 0)
OTlkYTczYWM.set_editor_property("desc", "No description provided")
OTlkYTczYWM.set_editor_property("default_value", 10.000000)
OTlkYTczYWM.set_editor_property("parameter_name", "offset")

# Inserting pins to custom expression
OTE4ZTdjMjg.set_editor_property("inputs", [NTY4NjZlYmM, MjIwMWRjMDk, NjQzYjMwMWI])

# Connecting input pins
unreal.MaterialEditingLibrary.connect_material_expressions(OWE1ZjRjZjk, "", OTE4ZTdjMjg, "size")
unreal.MaterialEditingLibrary.connect_material_expressions(MzM0YTc5MGM, "", OTE4ZTdjMjg, "direction")
unreal.MaterialEditingLibrary.connect_material_expressions(OTlkYTczYWM, "", OTE4ZTdjMjg, "offset")
