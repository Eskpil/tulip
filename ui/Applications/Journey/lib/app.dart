import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:journey/entities/bloc/entities_bloc.dart';
import 'package:journey/entities/bloc/entities_event.dart';
import 'package:journey/entities/view/entities.dart';
import 'package:journey/repositories/entity.dart';

class App extends StatelessWidget {

  final EntityRepository entityRepository = EntityRepository();

  @override
  Widget build(BuildContext context) {
    return MultiBlocProvider(
        providers: [
          BlocProvider(
            create: (_) => EntitiesBloc(
              entityRepository: entityRepository,
            )..add(EntitiesStarted()),
          ),
        ],
        child: MaterialApp(
          title: "Joruney",
          initialRoute: "/",
          routes: {
            "/": (_) => const EntitiesPage(),
          },
        ),
    );
  }
}