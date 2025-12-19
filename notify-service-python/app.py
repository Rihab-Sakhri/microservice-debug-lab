from flask import Flask, jsonify, request

app = Flask(__name__)

users_notified = []

@app.route("/notify", methods=["POST"])
def notify_user():
    data = request.get_json()
    user_id = data.get("id")

    # Intentional bug: appending user_id directly without validation
    if user_id == 0:
        return jsonify({"error": "Invalid user ID"}), 400

    users_notified.append(user_id)
    return jsonify({"message": f"User {user_id} notified!"})

@app.route("/notifications", methods=["GET"])
def get_notifications():
    return jsonify(users_notified)

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)