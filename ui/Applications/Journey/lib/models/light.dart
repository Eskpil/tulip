import 'package:journey/models/command.dart';
import 'package:journey/models/entity.dart';

import 'package:format/format.dart';
import 'dart:convert';
import 'package:http/http.dart' as http;

class BinaryLightSwitchCommand extends Subcommand {
  @override
  Map<String, dynamic> toJson() => {
    "name": "opposite",
  };
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