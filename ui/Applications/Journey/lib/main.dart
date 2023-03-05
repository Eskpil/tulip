import 'package:flutter/material.dart';
import 'package:journey/components/card.dart';
import 'package:journey/models/entity.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({super.key});

  // This widget is the root of your application.
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Journey',
      theme: ThemeData(
        primarySwatch: Colors.amber,
      ),
      home: const MyHomePage(),
    );
  }
}

class MyHomePage extends StatefulWidget {
  const MyHomePage({super.key});

  @override
  State<MyHomePage> createState() => _MyHomePageState();
}

class _MyHomePageState extends State<MyHomePage> {
  late Future<List<Entity>> futureEntities;

  @override
  void initState() {
    super.initState();
    futureEntities = fetchEntities();
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: Colors.black54,
      body: FutureBuilder<List<Entity>>(
        future: futureEntities,
        builder: (context, snapshot) {
          if (snapshot.hasData) {
            List<Widget> cards = [];

            for (final entity in snapshot.data!) {
              cards.add(
                Container(
                  margin: const EdgeInsets.all(18),
                  child: CardWidget(entity: entity,),
                ),
              );
            }

            return GridView.count(
              padding: const EdgeInsets.all(12),
              crossAxisCount: 3,
              children: cards,
            );
          } else if (snapshot.hasError) {
            return Text('${snapshot.error}');
          }

          // By default, show a loading spinner.
          return const CircularProgressIndicator();
        },
      ),
    );
  }
}
