import 'package:journey/models/command.dart';
import 'package:journey/models/entity.dart';

import 'dart:async';
import 'package:format/format.dart';
import 'dart:convert';
import 'package:http/http.dart' as http;
import 'package:sprintf/sprintf.dart';

class BinaryLightSwitchCommand extends Subcommand {
  @override
  Map<String, dynamic> toJson() => {
    "name": "opposite",
  };
}

class Light {
  static void opposite(Entity entity) async {
    final command = Command(subcommand: BinaryLightSwitchCommand());

    final uri = Uri.parse(format("http://localhost:8000/entities/%s/action/", entity.id));
    final response = await http
        .patch(uri, body: json.encode(command));

    if (response.statusCode == 200) {

    } else {
      final code = response.statusCode;
      throw Exception("Could not change light state to opposite: $code");
    }
  }
}