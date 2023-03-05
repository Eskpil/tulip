import 'package:journey/core/models/command.dart';
import 'package:journey/entities/models/entity.dart';

import 'package:format/format.dart';
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:journey/entity/models/entity_state.dart';

class BinaryLightSwitchCommand extends Subcommand {
  @override
  Map<String, dynamic> toJson() => {
    "name": "opposite",
  };
}

class LightState {
  final Object color;
  final String colorMode;
  final String state;

  LightState({
    required this.color,
    required this.colorMode,
    required this.state
  });

  factory LightState.fromParent(EntityState state) {
    return LightState(
      color: state.state["color"],
      colorMode: state.state["color_mode"],
      state: state.state["state"],
    );
  }
}


class Light {
  static void opposite(Entity entity) async {
    final command = Command(subject: "light", subcommand: BinaryLightSwitchCommand());

    final uri = Uri.parse(format("http://localhost:8000/entities/{}/action/", entity.id));
    final response = await http
        .patch(uri, body: json.encode(command), headers: { "Content-Type": "application/json"});

    if (response.statusCode == 200) {

    } else {
      final code = response.statusCode;
      throw Exception("Could not change light state to opposite: $code");
    }
  }
}